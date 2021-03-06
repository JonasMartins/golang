import { Menu, Divider, Text } from "@mantine/core";
import { Settings, Search, Photo, MessageCircle, Trash, Logout } from "tabler-icons-react";
import { useRouter } from "next/dist/client/router";
import { useLogoutMutation } from "@/generated/graphql";
//import withUrqlClientDef from "@/components/hoc/HocWithUrqlClient";

interface SettingsMenuProps {}

const SettingsMenu: React.FC<SettingsMenuProps> = () => {
	const router = useRouter();

	const [{}, logout] = useLogoutMutation();

	const HandleLogout = async () => {
		await logout();
		router.push("/login");
	};

	return (
		<Menu>
			<Menu.Label>Application</Menu.Label>
			<Menu.Item icon={<Settings size={14} />}>Settings</Menu.Item>
			<Menu.Item icon={<MessageCircle size={14} />}>Messages</Menu.Item>
			<Menu.Item icon={<Photo size={14} />}>Gallery</Menu.Item>
			<Menu.Item
				icon={<Search size={14} />}
				rightSection={
					<Text size="xs" color="dimmed">
						⌘K
					</Text>
				}
			>
				Search
			</Menu.Item>

			<Divider />

			<Menu.Label>Danger zone</Menu.Label>
			<Menu.Item icon={<Logout size={14} />} onClick={HandleLogout}>
				Logout
			</Menu.Item>
			<Menu.Item color="red" icon={<Trash size={14} />}>
				Delete my account
			</Menu.Item>
		</Menu>
	);
};

export default SettingsMenu;
