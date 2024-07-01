# Findings

Things that I have found that happen during a phase either of Build or Deployment.

## Build

N/A

## Deployment

You shouldn't use the `workflow_dispatch` from tag. You should use `workflow_dispatch` from a branch and take an input for the version that you want to deploy. If you run from a tag version, then the pipelines have to be in working condition in that version.

They will not use the latest version of the pipeline and just pull that release package.

