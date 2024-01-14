# Go: E-Commerce Ed.

Aplicação de E-Commerce de exemplo composta por várias aplicações separadas: uma interface de usuário (que fornece conteúdo ao usuário final como páginas da web); uma API de back end (chamada pela interface de usuário conforme necessário) e um microserviço que executa apenas uma tarefa (construção dinâmica de faturas PDF e envio para clientes como anexo de e-mail).

A aplicação venderá itens individuais e permitirá que os usuários adquiram uma assinatura mensal. Todas as transações com cartão de crédito serão processadas por meio do **Stripe**, um dos sistemas de processamento de pagamento mais populares hoje em dia, com uma rica API disponível em mais de 35 países e compatível com mais de 135 moedas.

### Terminal Virtual

- Construção de uma aplicação web funcional em Go para processamento seguro de transações "não presenciais" 
- Uso do pacote `html/template` para **renderizar a interface do usuário**.
- **Processamento de pagamentos** com cartão de crédito de forma segura, integrado com a **API do Stripe**.

### Aplicação Web

- Construção de um site que **permite** aos usuários **comprar produtos ou adquirir uma assinatura mensal.**
- Processamento de **compras individuais e assinaturas recorrentes**.
- Tratamento de _**cancelamentos**_, _**reembolsos**_ e _**armazenamento**_ de informações de transação em um banco de dados.
- Implementação de **funcionalidades de reembolso, cancelamento de assinatura, autenticação de sessão e tokens de autenticação**.
- Gerenciamento de usuários (**adicionar/editar/excluir**), **redefinição** segura **de senhas** e **logoff** instantâneo com **websockets**.

### Microserviço

- Desenvolvimento de um microserviço independente que **recebe um payload JSON de uma compra**.
- Geração de uma **fatura PDF** com informações do payload JSON.
- Criação de um **e-mail para o cliente com** anexo da fatura **PDF**.
- **Envio** do e-mail.

Todas essas partes (interface de usuário, back end e microserviço) são construídas usando um único código-fonte que produz múltiplos binários, facilitado pelo Gnu Make.