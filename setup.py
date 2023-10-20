from setuptools import setup, find_packages
from aws_sso_role_switcher import __version__

with open("README.md", "r") as fh:
    long_description = fh.read()

with open('requirements.txt') as f:
    requirements = f.read().splitlines()

setup(
    name='aws-sso-role-switcher',
    version=__version__,
    py_modules=['aws-sso-role-switcher'],
    author="Parker Barthlome",
    author_email="pbbarthlome@gmail.com",
    description="This is a CLI application to switch your role using autocompletion by parsing your config/config file and then performs an AWS SSO login.",
    license="MIT",
    long_description=long_description,
    long_description_content_type="text/markdown",
    download_url="https://github.com/Pacobart/aws-sso-role-switcher/archive/0.1.0.tar.gz",
    keywords=['AWS', 'ROLE', 'GIT', 'PROFILE', 'AUTOCOMPLETE'],
    packages=find_packages(),
    scripts=['bin/aws-sso-role-switcher.sh', 'bin/aws-sso-role-switcher'],
    install_requires=[
        'prompt_toolkit',
        'setuptools'
    ],
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ],
    python_requires='>=3.9'
)