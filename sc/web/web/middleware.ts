import type { NextRequest } from "next/server";
import { NextResponse } from "next/server";
import { COOKIE_NAME_SERVER, existentRoutes } from "./utils/consts";

export function middleware(req: NextRequest) {
	if (!COOKIE_NAME_SERVER) {
		return NextResponse.error();
	}

	let cookie = req.cookies.get(COOKIE_NAME_SERVER);
	if (!cookie && !req.nextUrl.pathname.startsWith("/login")) {
		if (existentRoutes.includes(req.nextUrl.pathname)) {
			return NextResponse.redirect(new URL("/login", req.url));
		}
	}

	if (req.nextUrl.pathname.startsWith("/logout")) {
		return NextResponse.redirect(new URL("/login", req.url));
	}

	return NextResponse.next();
}
