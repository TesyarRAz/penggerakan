"use client"
import { Button } from "@/components/ui/button";
import { useSession } from "next-auth/react";
import Image from "next/image";
import Link from "next/link";
import { useRouter } from "next/navigation";
import React from "react";
import deleteModul from "../_actions/delete-module-action";

function ModuleCard({ 
    module: {
        id,
        name,
        created_at,
    }
 }: {
    module: ModuleResponse
 }) {
  const { data: session } = useSession()
  const router = useRouter()

  const handleDelete = async () => {
    if (!session) return

    if (!confirm('Are you sure you want to delete this course?')) return

    const ok = await deleteModul(session, id)

    if (ok) {
      router.refresh()
    } else {
      alert('Failed to delete course')
    }
  }

  return (
    <div className="max-w-xs bg-white rounded-lg shadow-md overflow-hidden m-4">
      <picture>
        <img className="object-cover h-48 w-full" src={""} alt={id} />
      </picture>
      <div className="p-4">
        <h3 className="text-lg text-gray-400 font-bold mb-2">{name}</h3>
        <p className="text-gray-600 text-sm mb-4">{created_at}</p>
        <div className="flex items-center justify-end">
          <Button className="py-2.5 px-5 me-2 mb-2 text-sm font-medium text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-100 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700" asChild>
            <Link href={`/dashboard/modules/${id}`}>
              Module
            </Link>
          </Button>
          <Button className="py-2.5 px-5 me-2 mb-2 text-sm font-medium text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-100 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700" asChild>
            <Link href={`/dashboard/courses/${id}`}>
              Edit
            </Link>
          </Button>
          <Button type="button"
            className="focus:outline-none text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-red-600 dark:hover:bg-red-700 dark:focus:ring-red-900"
            onClick={handleDelete}
          >
            Hapus
          </Button>
        </div>
      </div>
    </div>
  );
}

export default ModuleCard