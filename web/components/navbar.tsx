"use client"

import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";
import { ScrollArea, ScrollBar } from "./ui/scroll-area";
import { cn } from "@/lib/utils";

const menus = [
  {
    name: "Course",
    href: "/courses",
  }
]

interface NavbarProps extends React.HTMLAttributes<HTMLDivElement> { }

export function Navbar({ className, ...props } : NavbarProps) {
  const pathname = usePathname()

  return (
    <div className="relative">
      <ScrollArea>
        <div className={cn("mb-4 flex items-center", className)} {...props}>
          {menus.map((menu, index) => (
            <Link
              href={menu.href}
              key={menu.href}
              className={cn(
                "flex h-7 items-center justify-center rounded-full px-4 text-center text-sm transition-colors hover:text-primary",
                pathname?.startsWith(menu.href) ||
                (index === 0 && pathname === "/")
                ? "bg-muted font-mediumm text-primary"
                : "text-muted-foreground"
              )}
            >
              {menu.name}
            </Link>
          ))}
        </div>
        <ScrollBar orientation="horizontal" className="invisible" />
      </ScrollArea>
    </div>
  );
}