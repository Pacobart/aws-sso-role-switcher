#!/usr/bin/env python3

import configparser
import os
import sys

from pathlib import Path
from prompt_toolkit import prompt
from prompt_toolkit.completion import FuzzyWordCompleter
from prompt_toolkit.validation import Validator
from . import __version__

REGIONS = [
    "af-south-1",
    "ap-east-1",
    "ap-northeast-2",
    "ap-northeast-1",
    "ap-south-1",
    "ap-south-2",
    "ap-southeast-1",
    "ap-southeast-2",
    "ap-southeast-3",
    "ap-southeast-4",
    "ca-central-1",
    "eu-central-1",
    "eu-north-1",
    "eu-west-1",
    "eu-west-2",
    "eu-west-3",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-1",
    "us-west-2"
]

class ASRS:

    def __init__(self):
        self.config = configparser.ConfigParser()
        default_path = os.path.join(Path.home(), '.aws/config')
        extended_path = os.environ.get('AWS_PROFILE_SWITCHER_PATH')
        self.version = f"{__version__}"
        if extended_path:
            path = extended_path
        else:
            path = default_path
        self.config.read(path)

    def run(self, sys_args):
        self.__init__()
        args = self.parse_arguments(sys_args)
        if args.version:
            print(f"Version: {self.version}")
            sys.exit(0)
        profile = self.set_aws_vars(args.profile)
        current_region = os.environ.get("AWS_DEFAULT_REGION", None)
        if current_region:
            if args.region:
                if args.region not in [current_region, current_region.replace("-","")]:
                    self.set_aws_region(args.region)
        else:
            self.set_aws_region(args.region)
        self.aws_set_credentials(profile)


    def set_aws_vars(self, arg):
        validator = Validator.from_callable(
            self.profile_validator,
            error_message='Not a valid profile name',
            move_cursor_to_end=True)
        if self.profile_validator(arg):
            profile = arg
        else:
            profile = prompt('Enter Profile: ',
                            default=arg,
                            completer=FuzzyWordCompleter(self.config.sections()),
                            complete_while_typing=True,
                            validator=validator)
            profile = profile.replace('profile ', '')
        print(f"export AWS_PROFILE={profile}")
        return profile

    def profile_validator(self, text):
        if text in self.config.sections():
            return True
        else:
            return False

    @staticmethod
    def set_aws_region(arg):
        region = prompt('AWS_DEFAULT_REGION Not Set. Choose Region: ', default=arg,
                        completer=FuzzyWordCompleter(REGIONS))
        print(f"export AWS_DEFAULT_REGION={region}")

    @staticmethod
    def region_validator(text):
        if text in REGIONS:
            return True
        else:
            return False

    @staticmethod
    def parse_arguments(sys_args):
        import argparse
        parser = argparse.ArgumentParser()
        parser.add_argument("profile", nargs='?', default="", help="Valid profile name found in your `.aws/config` file")
        parser.add_argument("region", nargs='?', default="", help="valid AWS Region")
        parser.add_argument("-v", "--version", help="Provide the version of CLI",
                            action="store_true")
        args = parser.parse_args()
        return args
    
    @staticmethod
    def aws_set_credentials(arg):
        os.system(f'aws configure export-credentials --profile {arg} --format env')
