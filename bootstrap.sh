    echo "##########################################################################"
    echo "###################### Updating packages ##############################"

    sudo apt-get update

    echo "##########################################################################"    
    echo "###################### Installing Git ##############################"

    sudo apt-get install git -y
   
    echo "##########################################################################"
    echo "############### Installing NodeJS on an Ubuntu Machine ###############"

    sudo curl -sL https://deb.nodesource.com/setup_8.x | sudo -E bash -

    sudo apt-get install nodejs -y

    echo "##########################################################################"
    echo "############# Installing and configuring Docker for Dev #######################"

    sudo apt-get install docker.io -y
    sudo usermod -G docker ubuntu
    sudo usermod -G docker vagrant
    docker --version

    sudo add-apt-repository \
    "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
    $(lsb_release -cs) \
    stable"

    sudo apt-get update

    sudo apt-get install docker-ce -y --force-yes


    #echo "########################################################################"
    #echo "###################### Installing Kubectl ##############################"

    #wget https://storage.googleapis.com/kubernetes-release/release/v1.10.1/bin/linux/amd64/kubectl && \
    #chmod +x kubectl && \
    #sudo mv kubectl /usr/local/bin/ 
    #kubectl version

