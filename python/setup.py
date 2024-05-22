from setuptools import  setup, find_packages

setup(
    name='cx-sdk',
    packages=find_packages(include=['cx_sdk']),
    version='0.1.0',
    description='An SDK for coralogix management apis',
    install_requires=[
        'grpcio',
        'grpcio-tools',
    ],
    author='Coralogix Ltd',
)
