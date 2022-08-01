import { Loader as LoaderMantine, Stack, Grid } from "@mantine/core";
import React from "react";

interface LoaderProps {}

const Loader: React.FC<LoaderProps> = ({}) => {
	return (
		<Grid>
			<Grid.Col span={4} />
			<Grid.Col span={4}>
				<Stack
					sx={() => ({ height: "100vh", justifyContent: "center", alignItems: "center" })}
				>
					<LoaderMantine />
				</Stack>
			</Grid.Col>
			<Grid.Col span={4} />
		</Grid>
	);
};

export default Loader;
