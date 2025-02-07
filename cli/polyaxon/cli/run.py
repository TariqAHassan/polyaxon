#!/usr/bin/python
#
# Copyright 2019 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# coding: utf-8
from __future__ import absolute_import, division, print_function

import sys

import click

from polyaxon.cli.check import check_polyaxonfile
from polyaxon.cli.getters.project import get_project_or_local
from polyaxon.exceptions import PolyaxonSchemaError
from polyaxon.logger import clean_outputs
from polyaxon.polyaxonfile import PolyaxonFile
from polyaxon.run.conda import run as conda_run
from polyaxon.run.docker import run as docker_run
from polyaxon.run.platform import run as platform_run
from polyaxon.specs import get_specification
from polyaxon.utils.formatting import Printer
from polyaxon.utils.validation import validate_tags


@click.command()
@click.option("--project", "-p", type=str)
@click.option(
    "-f",
    "--file",
    "polyaxonfile",
    multiple=True,
    type=click.Path(exists=True),
    help="The polyaxonfiles to run.",
)
@click.option(
    "--name",
    type=str,
    help="Name to give to this run, must be unique within the project, could be none.",
)
@click.option("--tags", type=str, help="Tags of this run, comma separated values.")
@click.option("--description", type=str, help="The description to give to this run.")
@click.option(
    "--upload",
    "-u",
    is_flag=True,
    default=False,
    help="To upload the repo before running.",
)
@click.option(
    "--log",
    "-l",
    is_flag=True,
    default=False,
    help="To start logging after scheduling the run.",
)
@click.option(
    "--local",
    is_flag=True,
    default=False,
    help="To start the run locally, with `docker` environment as default.",
)
@click.option("--conda_env", type=str, help="To start a local run with `conda`.")
@click.option(
    "--params",
    "-P",
    metavar="NAME=VALUE",
    multiple=True,
    help="A parameter to override the default params of the run, form -P name=value.",
)
@click.option("--profile", type=str, help="Name of profile to use for this run.")
@click.pass_context
@clean_outputs
def run(
    ctx,
    project,
    polyaxonfile,
    name,
    tags,
    description,
    upload,
    log,
    local,
    conda_env,
    params,
    profile,
):
    """Run polyaxonfile specification.

    Examples:

    \b
    ```bash
    $ polyaxon run -f file -f file_override ...
    ```

    Upload before running

    \b
    ```bash
    $ polyaxon run -f file -u
    ```

    Run and set description and tags for this run

    \b
    ```bash
    $ polyaxon run -f file -u --description="Description of the current run" --tags="foo, bar, moo"
    ```
    Run and set a unique name for this run

    \b
    ```bash
    polyaxon run --name=foo
    ```

    Run for a specific project

    \b
    ```bash
    $ polyaxon run -p project1 -f file.yaml
    ```

    Run with updated params

    \b
    ```bash
    $ polyaxon run -p project1 -f file.yaml -P param1=234.2 -P param2=relu
    ```
    """
    specification = check_polyaxonfile(
        polyaxonfile, params=params, profile=profile, log=False
    )

    owner, project_name = get_project_or_local(project)
    tags = validate_tags(tags)

    if local:
        try:
            run_spec = get_specification(specification.generate_run_data())
            run_spec.apply_context()
        except PolyaxonSchemaError:
            Printer.print_error(
                "Could not run this polyaxonfile locally, "
                "a context is required to resolve it dependencies."
            )
            sys.exit(1)
        if conda_env:
            conda_run(
                ctx=ctx,
                name=name,
                owner=owner,
                project_name=project_name,
                description=description,
                tags=tags,
                specification=run_spec,
                log=log,
                conda_env=conda_env,
            )
        else:
            docker_run(
                ctx=ctx,
                name=name,
                owner=owner,
                project_name=project_name,
                description=description,
                tags=tags,
                specification=run_spec,
                log=log,
            )
    else:
        platform_run(
            ctx=ctx,
            name=name,
            owner=owner,
            project_name=project_name,
            description=description,
            tags=tags,
            specification=specification,
            upload=upload,
            log=log,
            can_upload=all([upload, project]),
        )
