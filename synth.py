#!/usr/bin/env python3
import yaml
from jinja2 import Environment, FileSystemLoader

# load config file
with open("./kube/config.yml") as file:
    project = yaml.load(file, Loader=yaml.FullLoader)

# Initiate templates path
file_loader = FileSystemLoader("kube")
env = Environment(loader=file_loader)

# get a template
template = env.get_template("deploy-spec-template.yml")

# perform substitutions
output = template.render(env=project)

# save final file
file = open("./kube/deploy-spec.yml", "w")
file.write(output)
