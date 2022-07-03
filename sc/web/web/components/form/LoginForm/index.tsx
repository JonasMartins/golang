import type { NextPage } from "next";
import { useForm } from "@mantine/form";
import { TextInput, PasswordInput, Button, Group, Box, Text } from "@mantine/core";
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

	return (
		<form onSubmit={form.onSubmit(values => console.log(values))}>
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
					Submit
				</Button>
			</Group>
		</form>
	);
};

export default LoginForm;
