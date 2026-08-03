# Projeto Korp — Desafio Técnico

Ambiente completo em containers para o serviço `http-server-projeto-korp`, com proxy.
reverso NGINX, monitoramento via Prometheus/Grafana e provisionamento 100% automatizado
via Ansible.

## Estrutura do repositório

```
.
├── app/                          # Código-fonte Go + Dockerfile
│   ├── main.go
│   ├── go.mod
│   └── Dockerfile
├── nginx/
│   └── http-server-projeto-korp.conf
├── prometheus/
│   └── prometheus.yml
├── grafana/
│   ├── provisioning/
│   │   ├── datasources/datasource.yml
│   │   └── dashboards/dashboard.yml
│   └── dashboards/
│       └── http-server-projeto-korp-dashboard.json
├── ansible/
│   ├── inventory.ini
│   ├── playbook.yml
│   ├── requirements.yml
│   └── roles/
│       ├── docker/     # instala Docker + cria rede bridge
│       ├── deploy/     # build da imagem + docker compose up
│       └── validate/   # requisição HTTP de validação
└── docker-compose.yml
```

## Pré-requisitos

- Linux (testado em Linux Mint / base Ubuntu)
- Ansible instalado (`ansible --version`)
- Acesso à internet (para baixar imagens Docker e dependências Go)

## Como rodar tudo com um único comando

```bash
cd ansible
ansible-galaxy collection install -r requirements.yml
ansible-playbook -i inventory.ini playbook.yml --ask-become-pass
```

Isso vai:
1. Instalar o Docker (se necessário) e criar a rede bridge `korp-network`
2. Buildar a imagem do `http-server-projeto-korp`
3. Subir os 4 containers via Docker Compose (`http-server-projeto-korp`, `nginx`, `prometheus`, `grafana`)
4. Validar o serviço com uma requisição HTTP em `http://localhost/projeto-korp` e exibir a resposta no console

## Teste manual

```bash
curl http://localhost:80/projeto-korp
```

Resposta esperada:

```json
{"nome":"Projeto Korp","horario":"2026-08-02T12:00:00Z"}
```

## Acessos

| Serviço     | URL                          | Credenciais       |
|-------------|-------------------------------|--------------------|
| Aplicação   | http://localhost/projeto-korp | -                  |
| Prometheus  | http://localhost:9090         | -                  |
| Grafana     | http://localhost:3000         | admin / admin      |

O dashboard "Projeto Korp - http-server-projeto-korp" já vem provisionado automaticamente
no Grafana, com painéis de disponibilidade e volume de requisições.


