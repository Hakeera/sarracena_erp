# Sarracena ERP

Sarracena ERP é um sistema de gestão empresarial desenvolvido em Golang, utilizando o framework Gin e o banco de dados PostgreSQL. Este README documenta a estrutura do projeto, explicando a finalidade de cada diretório e os principais arquivos.

## Estrutura do Projeto

### **`main.go`** 🚀
Arquivo principal responsável por inicializar e executar a aplicação.

- **Principais Funções:**
  - **Carregar Variáveis de Ambiente:** Utiliza a biblioteca `godotenv` para carregar as variáveis definidas no arquivo `.env`.
  - **Conexão com Banco de Dados:** Configura e inicializa a conexão com o PostgreSQL, garantindo o fechamento adequado da pool de conexões.
  - **Inicializar Dependências:** Instancia os serviços e controladores necessários para o funcionamento da aplicação.
  - **Configurar Roteador:** Define as rotas utilizando o framework Gin.
  - **Iniciar Servidor:** Configura o servidor para escutar na porta `8080` e responde às requisições HTTP.

### **`configuration`** ⚙️
Contém configurações globais e utilitários essenciais para o funcionamento do projeto.

- **`logs/logger.go`**: Implementa logs estruturados e detalhados utilizando a biblioteca `zap`.
- **`validation/validation.go`**: Inclui funções para validação de erros de entrada do usuário.
- **`rest_err/rest_err.go`**: Define uma estrutura unificada para erros HTTP, incluindo código, mensagem e detalhes opcionais.

### **`controller`** 🧭
Gerencia a lógica de controladores, processando requisições HTTP e chamando os serviços correspondentes.

- **`routes/routes.go`**: Mapeia endpoints às funções do controlador e organiza as rotas da aplicação.
- **`user_controller.go`**: Define a interface do controlador e métodos para criar, buscar, atualizar e deletar usuários.
- **`create_user.go`**: Implementa o endpoint de criação de usuários, incluindo validação dos dados fornecidos.

### **`model`** 📊
Define as estruturas de dados e implementa a lógica de negócios do sistema.

- **`database/connection.go`**: Configura e gerencia a conexão com o banco de dados PostgreSQL.
  - **Funções:**
    - Inicializa a conexão com o banco.
    - Retorna a instância da pool de conexões.
    - Fecha a conexão quando o sistema é encerrado.
- **`request/user_request.go`**: Define a estrutura para validação de dados recebidos ao criar ou atualizar usuários.
- **`response/user_response.go`**: Estrutura os dados retornados ao cliente em respostas HTTP.
- **`service`**: Contém a lógica de negócio que conecta os controladores ao modelo.
  - Fornece métodos para criar, atualizar, buscar e deletar usuários.
- **`user_domain.go`**: Implementa o domínio do usuário, com métodos para manipular dados e realizar a criptografia de senhas.

### **`view`** 🖼️
Converte estruturas de domínio interno para respostas JSON retornadas ao cliente.

- **`convert_domain_to_response.go`**: Fornece a função `ConvertDomainToResponse`, que transforma um domínio de usuário em uma estrutura JSON para ser enviada ao cliente.
  - **Exemplo:** Mapeia campos como email, nome e idade para o formato esperado pelo cliente.

## Configuração do Ambiente 🛠️

1. **Variáveis de Ambiente:**
   - `DATABASE_URL`: URL de conexão com o banco de dados PostgreSQL.

2. **Instruções para Execução:**
   - Inicialize o banco de dados utilizando `InitDatabase()`.
   - Configure o servidor Gin e execute as rotas definidas.

3. **Dependências:**
   - [Gin](https://github.com/gin-gonic/gin)
   - [pgx](https://github.com/jackc/pgx)
   - [zap](https://github.com/uber-go/zap)

## Principais Funcionalidades 🌟

- **Gestão de Usuários:**
  - Criar, buscar por ID ou email, atualizar e deletar usuários.
- **Validação de Dados:**
  - Verifica campos obrigatórios, formato de email e regras de senha.
- **Segurança:**
  - Implementa criptografia de senha utilizando MD5.

## Como Contribuir 🤝

1. Faça um fork do repositório.
2. Crie uma branch para sua funcionalidade: `git checkout -b feature/nova-funcionalidade`.
3. Faça commit de suas alterações: `git commit -m 'Adiciona nova funcionalidade'`.
4. Envie para o repositório remoto: `git push origin feature/nova-funcionalidade`.
5. Abra um Pull Request.

## Licença 📄

Este projeto está licenciado sob a [Licença MIT](LICENSE).

