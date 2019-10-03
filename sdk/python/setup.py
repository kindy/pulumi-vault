# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import errno
from setuptools import setup, find_packages
from setuptools.command.install import install
from subprocess import check_call

class InstallPluginCommand(install):
    def run(self):
        install.run(self)
        try:
            check_call(['pulumi', 'plugin', 'install', 'resource', 'vault', '${PLUGIN_VERSION}'])
        except OSError as error:
            if error.errno == errno.ENOENT:
                print("""
                There was an error installing the vault resource provider plugin.
                It looks like `pulumi` is not installed on your system.
                Please visit https://pulumi.com/ to install the Pulumi CLI.
                You may try manually installing the plugin by running
                `pulumi plugin install resource vault ${PLUGIN_VERSION}`
                """)
            else:
                raise

def readme():
    with open('README.rst') as f:
        return f.read()

setup(name='pulumi_vault',
      version='${VERSION}',
      description='A Pulumi package for creating and managing vault cloud resources.',
      long_description=readme(),
      cmdclass={
          'install': InstallPluginCommand,
      },
      keywords='pulumi vault',
      url='https://pulumi.io',
      project_urls={
          'Repository': 'https://github.com/pulumi/pulumi-vault'
      },
      license='Apache-2.0',
      packages=find_packages(),
      install_requires=[
          'parver>=0.2.1',
          'pulumi>=1.0.0,<2.0.0',
          'semver>=2.8.1'
      ],
      zip_safe=False)
