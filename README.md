# Projeto Korp — Desafio Técnico

Ambiente completo em containers para o serviço `http-server-projeto-korp`, com proxy
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


# Evidências de Execução - Projeto Korp

  GNU nano 7.2                                                           evidencias/EVIDENCIAS.md *                                                                   
# Evidências de Execução — Projeto Korp

## Parte 1 - Organizar os Arquivos do Projeto
![Estrutura de Pastas do Projeto- Usado o comando ls-R](01-estrutura_pastas_projeto.png)

## Parte 2 — Serviço HTTP e Arquitetura
![Build da imagem](02-teste_servidor_GO.png) 
![Build da imagem](03-build_imagem_docker.png)
*Build da imagem Docker concluído com sucesso.*

## Parte 3 — Criação da Rede Docker em modo bridge
![Build da bridge](04-criacao_rede_bridge.png) )
*Build da network concluído com sucesso.*

![Containers em execução](05-docker_compose.png)
![Containers em execução](06-docker_compose.png)
![Containers em execução](07-docker_compose.png)
*Containers http-server-projeto-korp e nginx rodando via docker compose ps.*

## Parte 4 — NGINX
![Teste Ambiente via NGINX ](08-teste_nginx.png)
*Confirmação do proxy reverso funcional.*

## Parte 5 — Monitoramento
![Endpoint targets](09-terminal_metricas.png)
*Métricas de 'http_server_projeto_korp_requests_total' e 'http_server_projeto_korp_up'.*
![Prometheus targets](010-prometheus.png)
*Prometheus coletando métricas do serviço (status UP).*

![Dashboard Grafana](011-grafana.png)
![Dashboard Grafana](012-grafana.png)
*Dashboard com painéis de disponibilidade e volume de requisições.*

## Parte 6 — Ansible
![Remoção do Ambiente](013-remocao_containers_rede.png)
*Remoção dos containers e rede.*

# Parte 7 — Ansible
![Ansible playbook](015-playbook_ansible.png)
![Ansible playbook](016-playbook_ansible.png)
*Execução do playbook com um único comando, PLAY RECAP sem falhas.*

## Parte 8 — GitHub
![Publicação no GitHub](git_push_ok.png)
*Publicado o projeto no repositório.*


 
