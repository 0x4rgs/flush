# Flush Go (beta)
Flush Go é um streamer capaz de enviar os logs do sistema operacional para o Elasticsearch em tempo real.

# Instalação e Configuração

```bash
# Baixar o repositório
git clone https://github.com/0x4rgs/flush.git
```

```bash
# Ajustes no Linux/Mac
ulimit -n 99999
```

```bash
# Instalar dependências
go get github.com/joho/godotenv
```

```bash
# Configurar o ENV_LOCAL
LOG_PATH=/var/log/{FILE}.log
```

# Teste local

```bash
# Executando containers Elasticsearch/Kibana
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.11.0
docker run --link {ID_CONTAINER_ELASTICSEARCH}:elasticsearch -p 5601:5601 docker.elastic.co/kibana/kibana:7.11.0
```

# Iniciando o streamer

```bash
# Inicie o streaming de logs
$ go run main.go
```

# Acessando o Kibana

```bash
# Acesse a interface do kibana e configure o índice à partir do índice de exemplo "golang"
http://localhost:5601
```

## Observações

Dependendo da quantidade de dados contidos no arquivo de log configurado o elasticsearch em execução com docker não suportará o volume de requisições.

## Dúvidas e Sugestões

0x4rgs@protonmail.com

