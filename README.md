# 100-dodeva-31-demos-golang
GoLang based demos for Developping on AWS course 

## Demos

### Demo 1 : Introduction to the AWS SDK for GoLang
This demo demonstrates how simple it is to load the AWS SDK and start operating on services.

To load this demo: 

````bash
git checkout demo_1
````

To build the demo:

````bash
➜  100-dodeva-31-demos-golang git:(demo_1) ✗ cd demos/sdk
➜  sdk git:(demo_1) ✗ go build -o app.out
````

To run the demo:

````bash
➜  sdk git:(demo_1) ✗ ./app.out 
````

### AWS SDK Links

- [https://docs.aws.amazon.com/sdk-for-go/api/aws/](https://docs.aws.amazon.com/sdk-for-go/api/aws/)
    
    AWS SDK API foundations

- [https://docs.aws.amazon.com/sdk-for-go/api/aws/session/](https://docs.aws.amazon.com/sdk-for-go/api/aws/session/)

    Session API for configuring service clients

- [https://docs.aws.amazon.com/sdk-for-go/api/service/s3/](https://docs.aws.amazon.com/sdk-for-go/api/service/s3/)

    AWS SDK API for S3 Service

### Demo 2 : Use IAM and S3 together
This demo demonstrates how simple to use S3 and IAM services to accomplish some tasks. In its initial form, it creates a bucket following {prefix}-{account number}-{region} pattern

To load this demo: 

````bash
git checkout demo_2
````

To build the demo:

````bash
➜  100-dodeva-31-demos-golang git:(demo_2) ✗ cd demos/iam-demo
➜  sdk git:(demo_2) ✗ go build -o app.out
````

To run the demo:

````bash
➜  sdk git:(demo_2) ✗ ./app.out -prefix={some bucket name prefix}
````

### AWS SDK Links

- [https://docs.aws.amazon.com/sdk-for-go/api/aws/](https://docs.aws.amazon.com/sdk-for-go/api/aws/)
    
    AWS SDK API foundations

- [https://docs.aws.amazon.com/sdk-for-go/api/aws/session/](https://docs.aws.amazon.com/sdk-for-go/api/aws/session/)

    Session API for configuring service clients

- [https://docs.aws.amazon.com/sdk-for-go/api/service/s3/](https://docs.aws.amazon.com/sdk-for-go/api/service/s3/)

    AWS SDK API for S3 Service

- [https://docs.aws.amazon.com/sdk-for-go/api/service/s3/](https://docs.aws.amazon.com/sdk-for-go/api/service/iam/)

    AWS SDK API for IAM Service