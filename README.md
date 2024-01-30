# albconv
albconv is a command-line tool for converting AWS Application Load Balancer (ALB) access logs.

Supported formats.  
* JSON

Additionally, the tool is capable of handling multiple file inputs.  
It seamlessly concatenates and converts these files in the sequence they are specified.

![albconv demo GIF](docs/images/demo.gif)

For detailed information about the structure and contents of AWS Application Load Balancer access logs, refer to the official AWS documentation.  
[Access logs for your Application Load Balancer - AWS Documentation](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-access-logs.html)

## Install
### MacOS
```sh
brew install okkun-sh/tap/albconv
```

### Linux
```sh
wget https://github.com/okkun-sh/albconv/releases/download/v0.1.0/albconv_Linux_x86_64.tar.gz
tar zxvf albconv_Linux_x86_64.tar.gz
chmod +x ./albconv
sudo mv ./albconv /usr/local/bin/albconv
```

### Windows
Download from [Releases](https://github.com/okkun-sh/albconv/releases).

## Usage
```sh
albconv alb.log.gz alb2.log.gz
```

When used in conjunction with the [jq](https://github.com/jqlang/jq) command, the capabilities of albconv are greatly enhanced.
```sh
albconv alb.log.gz alb2.log.gz | jq .[].type
```

## License
albconv is licensed under the MIT License. See [LICENSE](LICENSE) for more information.
