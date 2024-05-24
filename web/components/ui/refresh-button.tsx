"use client"
import { useRouter } from "next/navigation"
import { useTransition } from "react"
import { Button } from "./button"
import { RiRefreshLine } from "react-icons/ri"

const RefreshButton = ({
    onRefresh
}: {
    onRefresh?: () => void
}) => {

    const [isPending, startTransition] = useTransition()
    const router = useRouter()

    const handleRefresh = () => {
        startTransition(() => {
            if (onRefresh) {
                onRefresh()
            } else {
                router.refresh()
            }
        })
    }

    return (
        <Button
            variant="blank"
            type="button"
            className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm  text-center inline-flex items-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800 p-3 mr-1"
            onClick={handleRefresh}
        >
            <RiRefreshLine />
        </Button>
    )
}

export default RefreshButton