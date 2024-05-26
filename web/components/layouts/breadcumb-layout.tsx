import React from 'react'
import RefreshButton from '../ui/refresh-button'

const BreadcumbLayout = ({
    title,
    children
}: {
    title?: string | null,
    children: React.ReactNode
}) => {
    return (
        <div>
            <div className="flex items-center justify-between my-3 py-3 px-3 rounded-lg bg-slate-600">
                <div>
                    <h4 className="font-sans font-semibold">{title}</h4>
                </div>
                <div className="justify-end">
                    <RefreshButton />
                </div>
            </div>

            {children}
        </div>
    )
}

export default BreadcumbLayout
