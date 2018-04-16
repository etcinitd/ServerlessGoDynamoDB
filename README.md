# ServerlessGoDynamoDB
# 1.Install go

$ tar -C /usr/local -xzf go1.10.1.linux-amd64.tar.gz

$ export PATH=$PATH:/usr/local/go/bin

$ export GOPATH=$HOME/go

$ export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

$ export GOBIN=$HOME/go/bin

# 2.Install nodejs

$ curl -sL https://deb.nodesource.com/setup_8.x | sudo -E bash -

$ sudo apt-get install -y nodejs

# 3.install serverless

$ npm install -g serverless

# 4.Create AWS account

https://aws.amazon.com/premiumsupport/knowledge-center/create-and-activate-aws-account/

# 5.Create an I AM user and configure credentials

https://serverless.com/framework/docs/providers/aws/guide/credentials/

$ serverless config credentials (-o) --provider aws --key XXX --secret YYY

# 6.Compile

Change into service directory and compile :

$ make

# 7.Deploy

$ serverless deploy (Or $ sls deploy)

# 8.Add Device

Enter data that you want to insert in /lib/data.json.

$ curl -X POST -H "Content-Type: application/json" -d @lib/data.json  https://XXX.execute-api.us-east-1.amazonaws.com/dev/devices

Note : put the url which is created when you deploy service.

# 9.Get Device By ID

$ curl https://XXX.execute-api.us-east-1.amazonaws.com/dev/devices/"id7"

