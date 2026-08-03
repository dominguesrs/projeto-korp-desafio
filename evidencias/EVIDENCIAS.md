                                                                   
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


 
