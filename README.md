# HMV

HMV é uma API de gestão do ciclo de vida emergências no Hospital Moinhos de Vento, envolvendo pacientes, socorrista e analistas do hospital.

## Requisitos:

- golang@1.17.1

## Rodando o projeto:

Executar o seguinte comando:

> make start

## Configurações padrão:

1) Servidor:
- server.address: :8080 // endereço padrão do servidor
- server.development_environment: true // modo de operação do servidor (true: desenvolvimento/false: release)

2) Logging:
- logging.development_environment: false // modo de logging (true: debug/false: produção)


A qualquer momento, é possível alterar as configurações da aplicação através do arquivo `settings.yaml`, que se encontra na raiz do projeto.

## Testando as APIs:

Para poder testar as APIs da aplicação, você pode importar a [collection do projeto](https://www.getpostman.com/collections/ab992d09a3eb6ff74a44) no Postman (no link, em formato JSON).
