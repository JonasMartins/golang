import { Menu, Divider, Text, ActionIcon } from "@mantine/core";
import {
	IconSettings,
	IconSearch,
	IconPhoto,
	IconMessageCircle,
	IconTrash,
	IconLogout,
	IconDots,
} from "@tabler/icons";
import { useRouter } from "next/dist/client/router";
import { useLogoutMutation } from "@/generated/graphql";
import { persistor } from "@/app";
import { clearState } from "@/features/chat/chatSlicer";
import { useDispatch } from "react-redux";
import { setOpenSettingsModal } from "@/features/layout/modalSlicer";

interface SettingsMenuProps {}

const SettingsMenu: React.FC<SettingsMenuProps> = () => {
	const router = useRouter();
	const dispatch = useDispatch();
	const [{}, logout] = useLogoutMutation();

	const HandleLogout = async () => {
		await logout();
		await persistor.purge();
		dispatch(clearState());
		router.push("/login");
	};

	return (
		<Menu shadow="md" transition="scale-y">
			<Menu.Target>
				<ActionIcon>
					<IconDots size={14} />
				</ActionIcon>
			</Menu.Target>

			<Menu.Dropdown>
				<Menu.Item
					icon={<IconSettings size={14} />}
					onClick={() => dispatch(setOpenSettingsModal(true))}
				>
					Settings
				</Menu.Item>
				<Menu.Item icon={<IconMessageCircle size={14} />}>Messages</Menu.Item>
				<Menu.Item icon={<IconPhoto size={14} />}>Gallery</Menu.Item>
				<Menu.Item
					icon={<IconSearch size={14} />}
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
				<Menu.Item icon={<IconLogout size={14} />} onClick={HandleLogout}>
					Logout
				</Menu.Item>
				<Menu.Item color="red" icon={<IconTrash size={14} />}>
					Delete my account
				</Menu.Item>
			</Menu.Dropdown>
		</Menu>
	);
};

export default SettingsMenu;
