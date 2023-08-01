package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{ //Cadastra um usuario
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},

	{ //Buscar todos usuarios
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: false,
	},

	{ // Busca um usuário por ID
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},

	{ // Atualiza um usuario por ID
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},

	{ //Deleta um usuario por ID
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: false,
	},
	/////////

}
