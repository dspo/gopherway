# go-micro introduce

## Step-1 install consul
0. [download consul binary file](https://www.consul.io/downloads.html) <-- click here;  
0. unzip package; make sure `consul` binary is available on the `PATH` ;  
0. open a terminal, verifying the installation.
    ```terminal
    $ consul
    usage: consul [--version] [--help] <command> [<args>]
    
    Available commands are:
        agent          Runs a Consul agent
        event          Fire a new event
    
    ```   
    not see this, be sure your `PATH` was set up properly.  
    
see more in [how to install consul](https://learn.hashicorp.com/consul/getting-started/install.html) <-- click here
    
## Step-2 run the consul agent
each datacenter must have at least one server;
 3 or 5 server is recommended.

0. starting the agent
    ```terminal
    $ consul agent -dev
    ==> Starting Consul agent...
    ==> Consul agent running!
               Version: 'v1.5.1'
               Node ID: '18339e2f-2b5e-cfa8-3ebe-a60cba24bb2a'
             Node name: 'Kaitlins-MBP'
            Datacenter: 'dc1' (Segment: '<all>')
                Server: true (Bootstrap: false)
           Client Addr: [127.0.0.1] (HTTP: 8500, HTTPS: -1, gRPC: 8502, DNS: 8600)
          Cluster Addr: 127.0.0.1 (LAN: 8301, WAN: 8302)
               Encrypt: Gossip: false, TLS-Outgoing: false, TLS-Incoming: false
               
    ... ...           
    ```
    you can see the agent running in server mode and has claimed leadership of the cluster.    

0. run `consul members` in another terminal
    ```terminal
    $ consul members
    Node          Address         Status  Type    Build  Protocol  DC   Segment
    Kaitlins-MBP  127.0.0.1:8301  alive   server  1.5.1  2         dc1  <all>
    ```

see more in [how to run the consul agent](https://learn.hashicorp.com/consul/getting-started/agent) <-- click here    

## Step-3 Writing a Go Service
> [see source here: writing a go service](https://micro.mu/docs/writing-a-go-service.html)  

1. create service proto
    1. make a proto3 script, define service
    0. generate the proto
0. write the service
    1. implements the interface defined for the handler
    0. initialises a micro.Service
    0. register the handler
    0. runs the service
## Step-4 Defining a client  
> [see details here](https://micro.mu/docs/go-micro.html#define-a-client)
1. creates a new micro.Service and initialises it
0. creates a new client uses NewXxxxService (implemented in xxx.micro.go)         
0. calls remote service
0. runs the client
