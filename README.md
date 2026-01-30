# Authenticate with Bitbucket OAuth

[![Step changelog](https://shields.io/github/v/release/bitrise-steplib/steps-authenticate-with-bitbucket-oauth?include_prereleases&label=changelog&color=blueviolet)](https://github.com/bitrise-steplib/steps-authenticate-with-bitbucket-oauth/releases)

Adds your Bitbucket OAuth configuration to the `.netrc` file.

<details>
<summary>Description</summary>

[This Step](https://github.com/bitrise-steplib/steps-authenticate-with-bitbucket-oauth) adds the authentication configuration (Bitbucket username and App password) to the `.netrc` file .
Please note that if you already have a `.netrc` file, the Step will create a backup of the original, and appends the configs to the current one.

### Configuring the Step
1. Add your **Bitbucket username**.
2. Add your Bitbucket **App Password**.

To get your Bitbucket App Password, follow the instructions below:
1. Log into your Bitbucket account.
2. In the left sidebar, click **App passwords**.
3. Click **Create app password**.
4. Give your password a descriptive label.
5. Select the permissions you'd like to grant to this token.
6. Click **Create**.

### Useful links
- [Learn more what the .netrc file format comprises of](https://everything.curl.dev/usingcurl/netrc#the-netrc-file-format)

### Related Steps
- [Activate SSH key (RSA private key)](https://www.bitrise.io/integrations/steps/activate-ssh-key)
- [Connect to OpenVPN Server](https://www.bitrise.io/integrations/steps/open-vpn)
</details>

## 🧩 Get started

Add this step directly to your workflow in the [Bitrise Workflow Editor](https://docs.bitrise.io/en/bitrise-ci/workflows-and-pipelines/steps/adding-steps-to-a-workflow.html).

You can also run this step directly with [Bitrise CLI](https://github.com/bitrise-io/bitrise).

## ⚙️ Configuration

<details>
<summary>Inputs</summary>

| Key | Description | Flags | Default |
| --- | --- | --- | --- |
| `username` | The username used for Bitbucket to login. | required, sensitive |  |
| `access_token` | To get your Bitbucket App Password, follow the instructions below:  1. Log into your Bitbucket account 2. In the upper-right corner of any page, click your profile photo, then click **Bitbucket Settings**. 3. In the left sidebar, click **App passwords**. 4. Click **Create app password**. 5. Give your password a descriptive label. 6. Select the permissions you'd like to grant to this token. 7. Click **Create**. | required, sensitive |  |
</details>

<details>
<summary>Outputs</summary>
There are no outputs defined in this step
</details>

## 🙋 Contributing

We welcome [pull requests](https://github.com/bitrise-steplib/steps-authenticate-with-bitbucket-oauth/pulls) and [issues](https://github.com/bitrise-steplib/steps-authenticate-with-bitbucket-oauth/issues) against this repository.

For pull requests, work on your changes in a forked repository and use the Bitrise CLI to [run step tests locally](https://docs.bitrise.io/en/bitrise-ci/bitrise-cli/running-your-first-local-build-with-the-cli.html).

Learn more about developing steps:

- [Create your own step](https://docs.bitrise.io/en/bitrise-ci/workflows-and-pipelines/developing-your-own-bitrise-step/developing-a-new-step.html)
