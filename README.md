Fazer download e Instalar o GO:
https://go.dev/dl/

Instalar GO no LINUX:
# Baixar o tar.gz
wget https://go.dev/dl/go1.24.4.linux-amd64.tar.gz

# Remover versões antigas (se existir)
sudo rm -rf /usr/local/go

# Extrair para /usr/local
sudo tar -C /usr/local -xzf go1.24.4.linux-amd64.tar.gz

# Adicionar Go ao PATH (se ainda não estiver)
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Verificar a instalação
go version


Criar Pasta go

Dentro da pasta go:
Criar pasta pkg: para pacotes go
Crar pasta src: para aplições
Criar pasta bin: para binarios/executaveis

Instalar o Rich Go Language no VSCode

Para compilar:
go build hello.go ou
go run hello.go