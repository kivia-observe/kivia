import { Card, CardContent } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";

function StatCardSkeleton() {
  return (
    <Card className="border shadow-sm">
      <CardContent className="pt-6">
        <Skeleton className="h-10 w-10 rounded-lg mb-4" />
        <Skeleton className="h-8 w-16 mb-1" />
        <Skeleton className="h-4 w-28 mb-1" />
        <Skeleton className="h-3 w-36" />
      </CardContent>
    </Card>
  );
}

export default function DashboardLoading() {
  return (
    <div className="p-8 w-full">
      <div className="mb-8">
        <Skeleton className="h-8 w-32 mb-2" />
        <Skeleton className="h-4 w-72" />
      </div>

      <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
        <StatCardSkeleton />
        <StatCardSkeleton />
      </div>

      <div className="mt-8">
        <div className="flex items-center justify-between mb-4">
          <Skeleton className="h-4 w-32" />
          <Skeleton className="h-4 w-12" />
        </div>
        <div className="rounded-xl border divide-y overflow-hidden bg-card">
          {[1, 2, 3].map((i) => (
            <div key={i} className="flex items-center gap-3 px-5 py-3.5">
              <Skeleton className="h-8 w-8 rounded-lg shrink-0" />
              <div className="flex-1">
                <Skeleton className="h-4 w-40 mb-1.5" />
                <Skeleton className="h-3 w-20" />
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
