System Requirements and Installation Guide
########################################

.. toctree::
  :maxdepth: 2

Installation
============
Following is how we build elephant node in ubuntu.

Ubuntu
******

Install git::

    $ sudo apt-get install -y git

Install essentials and Go ::

    $ sudo apt-get install -y software-properties-common
    $ sudo add-apt-repository -y ppa:gophers/archive
    $ sudo apt update
    $ sudo apt-get install -y golang-1.12-go

Setup basic workspace,In this instruction we use ~/dev/src/github.com/elastos as our working directory. If you clone the source code to a different directory, please make sure you change other environment variables accordingly (not recommended)::

    $ mkdir -p ~/dev/bin
    $ mkdir -p ~/dev/src/github.com/elastos/

Set correct environment variables::

    export GOROOT=/usr/local/opt/go@1.12/libexec
    export GOPATH=$HOME/dev
    export GOBIN=$GOPATH/bin
    export PATH=$GOROOT/bin:$PATH
    export PATH=$GOBIN:$PATH

Clone source code to $GOPATH/src/github/elastos folder,Make sure you are in the folder of $GOPATH/src/github.com/elastos::

    $ git clone https://github.com/elastos/Elastos.ELA.Elephant.Node.git

If clone works successfully, you should see folder structure like $GOPATH/src/github.com/elastos/Elastos.ELA.Elephant.Node/Makefile

Install Glide, Glide is a package manager for Golang. We use Glide to install dependent packages::

    $ cd ~/dev
    $ curl https://glide.sh/get | sh

Make,Build the node::

    $ cd $GOPATH/src/github.com/elastos/Elastos.ELA.Elephant.Node
    $ glide update & glide install (for chinese users you may need to use vpn to download dependency)

update config.json similar to the following content::

    {
      "Configuration": {
        "ActiveNet": "mainnet",
        "HttpInfoPort": 20333,
        "HttpInfoStart": true,
        "HttpRestPort": 20334,
        "HttpRestStart": true,
        "HttpWsPort": 20335,
        "HttpWsStart":true,
        "HttpJsonPort": 20336,
        "EnableRPC": true,
        "NodePort": 20338,
        "PrintLevel": 1,
        "PowConfiguration": {
          "PayToAddr": "EeEkSiRMZqg5rd9a2yPaWnvdPcikFtsrjE",
          "MinerInfo": "ELA",
          "MinTxFee": 100
        },
        "RpcConfiguration": {
          "User": "clark",
          "Pass": "123456",
          "WhiteIPList": [
            "127.0.0.1"
          ]
        }
      }
    }

Extra feature configure::

    {
      //Get CMC Apikey go to https://coinmarketcap.com/api/
      "Cmc":{
        "ApiKey":["5c8bcab7-a811-428d-9c2b-3f0326de4f66","237bd580-3c68-42f8-b9e8-fd201c4933ac"],
        // update schedule period, `m` stand for minutes
        "Inteval":"15m"
      }
    }


If you did not see any error message, congratulations, you have made the Elephant full node.

Run the node::

    $ sudo apt-get install build-essential
    $ make
    $ ./elephant

Nginx config Example::

    server { # simple reverse-proxy
        server_name  exmaple.com;
        access_log   /var/log/nginx/node.access.log;

        # pass requests for dynamic content to rails/turbogears/zope, et al

        location ~ ^/api/1/balance {
            proxy_pass http://localhost:20334;
            proxy_connect_timeout 120s;
            proxy_read_timeout 120s;
            proxy_send_timeout 120s;
            proxy_set_header Host $host;
            proxy_set_header X-Forwarded-For $remote_addr;
            proxy_set_header X-Real-IP $remote_addr;
        }

        location ~ ^/api/1/cmc {
            proxy_pass http://localhost:20334;
            proxy_connect_timeout 120s;
            proxy_read_timeout 120s;
            proxy_send_timeout 120s;
            proxy_set_header Host $host;
            proxy_set_header X-Forwarded-For $remote_addr;
            proxy_set_header X-Real-IP $remote_addr;
        }

        location ~ ^/api/1/createTx {
            proxy_pass http://localhost:20334;
            proxy_connect_timeout 120s;
            proxy_read_timeout 120s;
            proxy_send_timeout 120s;
            proxy_set_header Host $host;
            proxy_set_header X-Forwarded-For $remote_addr;
            proxy_set_header X-Real-IP $remote_addr;
        }

        location ~ ^/api/1/createVoteTx {
            proxy_pass http://localhost:20334;
            proxy_connect_timeout 120s;
            proxy_read_timeout 120s;
            proxy_send_timeout 120s;
            proxy_set_header Host $host;
            proxy_set_header X-Forwarded-For $remote_addr;
            proxy_set_header X-Real-IP $remote_addr;
        }

        location ~ ^/api/1/history {
            proxy_pass http://localhost:20334;
            proxy_connect_timeout 120s;
            proxy_read_timeout 120s;
            proxy_send_timeout 120s;
            proxy_set_header Host $host;
            proxy_set_header X-Forwarded-For $remote_addr;
            proxy_set_header X-Real-IP $remote_addr;
        }

        location ~ ^/api/1/sendRawTx {
            proxy_pass http://localhost:20334;
            proxy_connect_timeout 120s;
            proxy_read_timeout 120s;
            proxy_send_timeout 120s;
            proxy_set_header Host $host;
            proxy_set_header X-Forwarded-For $remote_addr;
            proxy_set_header X-Real-IP $remote_addr;
        }

        location ~ ^/api/1/pubkey {
            proxy_pass http://localhost:20334;
            proxy_connect_timeout 120s;
            proxy_read_timeout 120s;
            proxy_send_timeout 120s;
            proxy_set_header Host $host;
            proxy_set_header X-Forwarded-For $remote_addr;
            proxy_set_header X-Real-IP $remote_addr;
        }

        location ~ ^/api/1/currHeight {
            proxy_pass http://localhost:20334;
            proxy_connect_timeout 120s;
            proxy_read_timeout 120s;
            proxy_send_timeout 120s;
            proxy_set_header Host $host;
            proxy_set_header X-Forwarded-For $remote_addr;
            proxy_set_header X-Real-IP $remote_addr;
        }

        location ~ ^/api/1/dpos {
            proxy_pass http://localhost:20334;
            proxy_connect_timeout 120s;
            proxy_read_timeout 120s;
            proxy_send_timeout 120s;
            proxy_set_header Host $host;
            proxy_set_header X-Forwarded-For $remote_addr;
            proxy_set_header X-Real-IP $remote_addr;
        }

        listen 443 ssl; # managed by Certbot
        ssl_certificate /etc/letsencrypt/live/exmaple.com/fullchain.pem; # managed by Certbot
        ssl_certificate_key /etc/letsencrypt/live/exmaple.com/privkey.pem; # managed by Certbot
        include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
        ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot

    }

    server {
        if ($host = exmaple.com) {
            return 301 https://$host$request_uri;
        } # managed by Certbot


        server_name  exmaple.com;
        listen 80;
        return 404; # managed by Certbot
    }