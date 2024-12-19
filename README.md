# Terragrunt Reference Architecture 🚀

This repository serves as a comprehensive reference architecture for managing infrastructure using Terragrunt. It provides a structured approach to organizing and deploying your infrastructure as code.

## 🗂️ Project Structure

- `infra/`: Contains all infrastructure-related code.
  - `terraform/`: Houses Terraform modules.
  - `terragrunt/`: Contains Terragrunt configurations.
    - `_ENVS/`: Environment-specific configurations.
    - `_shared/`: Shared components and configurations.
    - `_templates/`: Templates for Terraform files.
    - `stack-*`: Example stacks.

## ⚙️ Key Concepts

- **Environments**: Configurations are managed per environment (e.g., `local`, `dev`, `prod`).
- **Stacks**: Infrastructure is organized into logical stacks (e.g., `stack-dx`, `stack-landing-zone`, `stack-webapp`).
- **Components**: Reusable infrastructure components are defined in `_shared/_components`.
- **Layers**: Each stack is composed of layers, allowing for modular and composable infrastructure.

## 🚀 Getting Started

1.  Clone the repository.
2.  Navigate to the `infra/terragrunt` directory.
3.  Configure your environment in `_ENVS/`.
4.  Deploy your infrastructure using Terragrunt.

## 🤝 Contributing

Feel free to contribute by opening issues or submitting pull requests.

## 📄 License

MIT License

## 🔮 Roadmap

- [ ] Add detailed documentation for each stack.
- [ ] Implement CI/CD pipelines.
- [ ] Include more example stacks.

## 💬 Support

Open an issue in the GitHub repository for any questions or problems.
