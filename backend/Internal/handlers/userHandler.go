package handlers

import (
	"ApiLogin/Internal/auth"
	"ApiLogin/Internal/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	DB *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{DB: db}
}

type cadastroRequest struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func (h *UserHandler) Cadastro(w http.ResponseWriter, r *http.Request) {
	var req cadastroRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "corpo da requisição inválido", http.StatusBadRequest)
		return
	}

	if req.Nome == "" || req.Email == "" || req.Senha == "" {
		http.Error(w, "nome, email e senha são obrigatórios", http.StatusBadRequest)
		return
	}

	hash, err := auth.HashPassword(req.Senha)
	if err != nil {
		http.Error(w, "erro ao processar senha", http.StatusInternalServerError)
		return
	}

	result, err := h.DB.ExecContext(r.Context(),
		"INSERT INTO usuarios (nome, email, senha) VALUES (?, ?, ?)",
		req.Nome, req.Email, hash)

	if err != nil {
		http.Error(w, "erro ao criar usuario", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"id":    id,
		"nome":  req.Nome,
		"email": req.Email,
	})
}

type loginRequest struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "corpo da requisição inválido", http.StatusBadRequest)
		return
	}

	var user models.User
	err := h.DB.QueryRowContext(r.Context(),
		"SELECT id_usuario, nome, email, senha FROM usuarios WHERE email = ?", req.Email,
	).Scan(&user.Id, &user.Nome, &user.Email, &user.Senha)

	if err == sql.ErrNoRows {
		http.Error(w, "credenciais inválidas", http.StatusUnauthorized)
		return
	}
	if err != nil {
		http.Error(w, "erro interno", http.StatusInternalServerError)
		return
	}

	if err := auth.VerificarSenha(req.Senha, user.Senha); err != nil {
		http.Error(w, "credenciais inválidas", http.StatusUnauthorized)
		return
	}

	token, err := auth.GerarToken(user.Id)
	if err != nil {
		http.Error(w, "erro ao gerar token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

type userResponse struct {
	Id    int64  `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func (h *UserHandler) BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.QueryContext(r.Context(), "SELECT id_usuario, nome, email FROM usuarios")
	if err != nil {
		http.Error(w, "erro ao buscar usuarios", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var usuarios []userResponse
	for rows.Next() {
		var u userResponse
		if err := rows.Scan(&u.Id, &u.Nome, &u.Email); err != nil {
			http.Error(w, "erro ao processar usuarios", http.StatusInternalServerError)
			return
		}
		usuarios = append(usuarios, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}
