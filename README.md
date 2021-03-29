# Golang!

Repositório que tem como objetivo registrar os meus estudos sobre Golang.


## Como escrever código Go

Essa é a minha tradução e versão do documento [How to Write Go Code](https://golang.org/doc/code) disponível no próprio site da linguagem.

### Organização de código 

Programas em Go são organizados em **pacotes**. Um *pacote* é uma coleção de "source files" no mesmo diretório que são compilados juntos. Funções, tipos, variáveis e constantes definidas em um source file são visíveis para todos os outros source files dentro do mesmo pacote. 

Um **repositório** contém um ou mais *módulos*. Um **módulo** é uma coleção de **pacotes** Go que são lançados juntos. Um repositório Go *normalmente contém apenas um módulo*, localizado na raiz (root) do repositório. Um arquivo chamado *go.mod* declara o caminho do módulo *(module path*): O prefixo de importação para todos os pacotes dentro do módulo. O módulo contém os pacotes no diretório contendo o seu arquivo go.mod assim como os subdiretórios daquele diretório, até o próximo subdiretório contendo outro arquivo go.mod (se tiver).

Note que você não precisa publicar seu código para um repositório remoto antes de conseguir efetuar o build. **Um módulo pode ser definido localmente sem pertencer a um repositório**. Entretanto, é um bom hábito organizar seu código como se você fosse publica-lo algum dia.

Cada caminho do módulo (*module's path*) não serve apenas como um prefixo de caminho de importação para os pacotes, mas também indica onde o comando go deve olhar para efetuar o download dele. Por exemplo, para efetuar o download do módulo golang.org/x/tools, the go command would consult the repository indicated by https://golang.org/x/tools.

Um caminho de importação (*import path*) é uma string usada para importar um pacote. Um caminho de importação de um pacote é o seu caminho de importação mais seu subdiretório dentro do módulo. Por exemplo, o módulo **github.com/google/go-cmp** contém um pacote no diretório **cmp/**. O caminho de importação do pacote é **github.com/google/go-cmp/cmp**. Pacotes na biblioteca padrão não tem um prefixo do caminho de módulo.

### Seu primeiro programa

Para compilar e rodar um programa simples, primeiro escolha um caminho de módulo (module path), nesse caso, vamos usar example.com/user/hello e criar o arquivo go.mod que o declara.

```bash
$ mkdir hello
$ cd hello
$ go mod init example.com/user/hello
```

O primeiro argumento em um source file Go deve ser o nome do pacote (*package name*). Comandos executáveis devem sempre usar *package main*.

Após isso, crie um arquivo chamado hello.go dentro do diretório contendo o seguinte código Go:

```Go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")
}
```


Agora você pode fazer o build e instalar essa aplicação com a CLI do Golang:

```bash
$ go install example.com/user/hello
```

Esse comando faz o build do comando hello, produzindo um executável binário. Então, instala o binário em *$HOME/go/bin/hello* (ou, no Windows, *%USERPROFILE%\go\bin\hello.exe*).

O diretório de instalação é controlado pelas variávies de ambiente **GOPATH** e **GOBIN**. Se GOBIN tiver algum valor, os binários são instalados no subdiretório bin do primeiro diretório na lista de GOPATH. Caso GOBIN não tenha valor, os binários são instalados no subdiretório bin no valor padrão de GOPATH (*$HOME /go* ou *%USERPROFILE%\go*).

Você pode usar o comando go env para inserir manualmente o valor de variáveis de ambiente:

```bash
$ go env -w GOBIN=/somewhere/else/bin
```

Para remover o valor de uma variável de ambiente inserida anteriormente utilizando *go env -w* use *go env -u*:

```bash
$ go env -u GOBIN
```

Comandos como *go install* aplicam-se dentro do contexto do módulo contendo o diretório de trabalho atual. Se o diretório atual não está dentro do módulo *example/com/user/hello*, *go install* pode falhar.

Para conveniência, comandos go aceitam caminhos relativos ao diretório de trabalho, e para o pacote padrão no diretório atual se nenhum outro caminho é passado. Desse jeito, no nosso diretório de trabalho, todos os comandos abaixo são válidos

```bash
$ go install example.com/user/hello
```
```bash
$ go install .
```
```bash
$ go install 
```

Agora, vamos rodar nosso programa para garantir que funciona. 

```bash
# Caso esteja no Windows, verificar o link: https://github.com/golang/go/wiki/SettingGOPATH
# for setting %PATH%.
$ export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
$ hello
Hello, world.
$
```

O comando go localiza o repositório que contém um determinado caminho de módulo, solicitando uma URL HTTPS correspondente e lendo metadados incorporados na resposta HTML (consulte *go help importpath*). Muitos serviços de hospedagem já fornecem esses metadados para repositórios que contêm código Go, portanto, a maneira mais fácil de tornar seu módulo disponível para outros usarem é geralmente fazer com que o caminho do módulo corresponda à URL do repositório.
