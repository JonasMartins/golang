import { LoginInput, useLoginMutation } from "@/generated/graphql";
import { Button, Group, PasswordInput, TextInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import type { NextPage } from "next";
import { IconEyeCheck, IconEyeOff } from "@tabler/icons";
import { useRouter } from "next/router";
import { useState } from "react";
import { setToken } from "@/features/token/tokenSlice";
import { useDispatch } from "react-redux";

type input = {
	email: string;
	password: string;
};

const LoginForm: NextPage = () => {
	const router = useRouter();
	const dispatch = useDispatch();

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

	const [{}, login] = useLoginMutation();

	const [input, setInput] = useState<input>({
		email: "",
		password: "",
	});

	const HandleLoginForm = async (values: LoginInput) => {
		const response = await login({ input: values });

		if (response.data?.login.errors.length) {
			const err = response.data?.login.errors[0];
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
			dispatch(setToken(response.data?.login.token!));
			router.push("/");
		}
	};

	return (
		<form onSubmit={form.onSubmit(values => HandleLoginForm(values))}>
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
				visibilityToggleIcon={({ reveal }) => (reveal ? <IconEyeOff /> : <IconEyeCheck />)}
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
