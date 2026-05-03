import { Skeleton } from "@/components/ui/skeleton";

export default function ProjectDetailLoading() {
  return (
    <div className="p-8 w-full">
      {/* Header */}
      <div className="mb-6">
        <Skeleton className="h-8 w-20 mb-4 rounded-md" />
        <Skeleton className="h-7 w-48 mb-1.5" />
        <Skeleton className="h-4 w-32" />
      </div>

      {/* Tabs */}
      <div className="mb-6 flex gap-1 border-b pb-0">
        <Skeleton className="h-10 w-28 rounded-none rounded-t-md" />
        <Skeleton className="h-10 w-32 rounded-none rounded-t-md" />
      </div>

      {/* Tab content — API Keys skeleton */}
      <div>
        <div className="flex items-center justify-between mb-6">
          <div>
            <Skeleton className="h-5 w-20 mb-1.5" />
            <Skeleton className="h-4 w-56" />
          </div>
          <Skeleton className="h-8 w-24 rounded-md" />
        </div>

        <div className="rounded-xl border overflow-hidden bg-card">
          <div className="bg-muted/30 px-4 py-3 flex gap-8">
            <Skeleton className="h-3.5 w-12" />
            <Skeleton className="h-3.5 w-12" />
            <Skeleton className="h-3.5 w-16" />
          </div>
          {[1, 2, 3].map((i) => (
            <div key={i} className="flex items-center gap-8 px-4 py-4 border-t">
              <Skeleton className="h-4 w-32" />
              <Skeleton className="h-5 w-14 rounded-full" />
              <Skeleton className="h-4 w-24" />
              <Skeleton className="h-7 w-16 ml-auto rounded-md" />
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
