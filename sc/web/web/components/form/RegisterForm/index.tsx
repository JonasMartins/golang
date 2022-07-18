import { RegisterUserInput, useRegisterUserMutation } from "@/generated/graphql";
import { Button, Group, PasswordInput, TextInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import type { NextPage } from "next";
import { EyeCheck, EyeOff } from "tabler-icons-react";
import { useRouter } from "next/router";
import { useState } from "react";

type input = {
	email: string;
	password: string;
};

const RegisterForm: NextPage = () => {
	const router = useRouter();

	const form = useForm({
		initialValues: {
			name: "",
			email: "",
			password: "",
			confirmPassword: "",
		},
		validate: {
			email: value => (/^\S+@\S+$/.test(value) ? null : "Invalid email"),
			password: value => (value.length >= 6 ? null : "Length Must be greather or equal to 6"),
			confirmPassword: (value, values) =>
				value !== values.password ? "Passwords did not match" : null,
		},
	});

	const [{}, register] = useRegisterUserMutation();

	const [input, setInput] = useState<input>({
		email: "",
		password: "",
	});

	const HandleRegisterForm = async (values: RegisterUserInput) => {
		const response = await register({
			input: {
				name: values.name,
				email: values.email,
				password: values.password,
			},
		});
		console.log(response);
		if (response.data?.registerUser.errors.length) {
			const err = response.data?.registerUser.errors[0];
			switch (err.field) {
				case "email":
					setInput(prevInput => ({
						...prevInput,
						email: err.message,
						password: "",
					}));
					break;
				case "password":
					setInput(prevInput => ({
						...prevInput,
						email: "",
						password: err.message,
					}));
					break;
				default:
					setInput(prevInput => ({
						...prevInput,
						email: "",
						password: "",
					}));
			}
		} else {
			router.push("/");
		}
	};

	return (
		<form onSubmit={form.onSubmit(values => HandleRegisterForm(values))}>
			<TextInput
				required
				label="Name"
				placeholder="Your Name"
				{...form.getInputProps("name")}
			/>

			<TextInput
				required
				label="Email"
				placeholder="your@email.com"
				{...form.getInputProps("email")}
				error={input.email}
			/>

			<PasswordInput
				placeholder="Password"
				required
				label="Password"
				error={input.password}
				{...form.getInputProps("password")}
				visibilityToggleIcon={({ reveal }) => (reveal ? <EyeOff /> : <EyeCheck />)}
			/>

			<PasswordInput
				placeholder="Confirm Password"
				required
				label="Confirm Password"
				{...form.getInputProps("confirmPassword")}
				visibilityToggleIcon={({ reveal }) => (reveal ? <EyeOff /> : <EyeCheck />)}
			/>
			<Group grow={true} mt="md">
				<Button variant="gradient" gradient={{ from: "indigo", to: "cyan" }} type="submit">
					Register
				</Button>
			</Group>
		</form>
	);
};

export default RegisterForm;
