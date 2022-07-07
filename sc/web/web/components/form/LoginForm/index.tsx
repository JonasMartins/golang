import { LoginInput, useLoginMutation } from "@/generated/graphql";
import { Button, Group, PasswordInput, TextInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import type { NextPage } from "next";
import { EyeCheck, EyeOff } from "tabler-icons-react";

const LoginForm: NextPage = () => {
	const form = useForm({
		initialValues: {
			email: "",
			password: "",
		},

		validate: {
			email: value => (/^\S+@\S+$/.test(value) ? null : "Invalid email"),
			password: value => (value.length >= 6 ? null : "Length Must be greather or equal to 6"),
		},
	});

	const [loginResult, login] = useLoginMutation();

	const HandleLoginForm = async (values: LoginInput) => {
		const response = await login({ input: values });
		console.log(response);
	};

	return (
		<form onSubmit={form.onSubmit(values => HandleLoginForm(values))}>
			<TextInput
				required
				label="Email"
				placeholder="your@email.com"
				{...form.getInputProps("email")}
			/>

			<PasswordInput
				placeholder="Password"
				required
				label="Password"
				{...form.getInputProps("password")}
				visibilityToggleIcon={({ reveal }) => (reveal ? <EyeOff /> : <EyeCheck />)}
			/>
			<Group grow={true} mt="md">
				<Button variant="gradient" gradient={{ from: "indigo", to: "cyan" }} type="submit">
					Login
				</Button>
			</Group>
		</form>
	);
};

export default LoginForm;
