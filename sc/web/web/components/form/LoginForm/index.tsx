import { LoginInput, useLoginMutation } from "@/generated/graphql";
import { Button, Group, PasswordInput, TextInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import type { NextPage } from "next";
import { EyeCheck, EyeOff } from "tabler-icons-react";
import { SERVER_URL } from "@/utils/consts";
import { withUrqlClient } from "next-urql";
import { dedupExchange, cacheExchange, fetchExchange } from "@urql/core";

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

		if (response.data?.login.token) {
			console.log(response);
		}
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

export default withUrqlClient(ssrExchange => ({
	url: SERVER_URL,
	exchanges: [dedupExchange, cacheExchange, ssrExchange, fetchExchange],
}))(LoginForm);
