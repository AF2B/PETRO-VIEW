## PETRO VIEW
### API RESTful para gerenciar leituras de sensores e dispositivos IoT.

### Recursos

* Coleta e armazenamento de leituras de sensores em tempo real.
* Consultas de histórico de leituras com filtros e paginação.
* Alertas e notificações para leituras fora do intervalo normal.
* Autenticação e autorização para acesso seguro à API.
* TODO

### Instalação
<code>
    git clone https://github.com/AF2B/PETRO-VIEW.git
    cd PETRO-VIEW

    go mod vendor
    go build

    ./PETRO-VIEW
</code>

### Uso Básico
<code>
    make run <br>
    make build <br>
    make docker-up <br>
    make docker-down <br>
    make test <br>
</code>

### Testes
Os testes unitários são escritos com o framework testing e os assert do testify. <br>
Para executar os testes, execute o seguinte comando: <br>
<code>
    go test ./...
</code>

### Informações de Contato:
- E-mail: andre.borbaaf2b@gmail.com
- Issues: andre.borbaaf2b@gmail.com