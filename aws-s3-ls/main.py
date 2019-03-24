#!/usr/bin/env python

from os import environ
import sys
import argparse
import boto3
import botocore
parser = argparse.ArgumentParser(description="""
    A tool to list an s3 bucket's public keys
""")
parser.add_argument("bucket")
parser.add_argument('--verbose', dest='verbose', action='store_true')
parser.set_defaults(verbose=False)
args = parser.parse_args()


def print_verbose(msg):
    print('--  {}'.format(msg))


def check_env():
    ok = True
    required_envs = ['AWS_ACCESS_KEY_ID', 'AWS_SECRET_ACCESS_KEY']
    for renv in required_envs:
        if environ.get(renv) is None:
            print('error: check_env: env is missing: {}'.format(renv))
            ok = False

    if not ok:
        print('info: exiting due to error')
        sys.exit(1)


def check_objects(bucket):
    client = boto3.client('s3')
    s3 = boto3.resource('s3')
    b = s3.Bucket(bucket)

    try:
        for obj in b.objects.all():
            if args.verbose:
                print_verbose(obj.key)
            c = client.get_object_acl(
                    Bucket=bucket,
                    Key=obj.key
            )
            for grant in c['Grants']:
                if args.verbose:
                    print_verbose(grant)
                public = True
                if ('URI' in grant['Grantee']) and (grant['Grantee']['URI'] == 'http://acs.amazonaws.com/groups/global/AllUsers'):
                    public = True
                else:
                    public = False

                if public:
                    print('{}'.format(obj.key))
    except botocore.exceptions.ClientError as e:
        print('error: AWS client exception: {}'.format(e))


def main():
    check_env()
    check_objects(args.bucket)
    if args.verbose:
        print_verbose('info: exiting normally')
    sys.exit(0)


if __name__ == '__main__':
    main()
