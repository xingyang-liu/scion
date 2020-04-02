@0xe6c88f91b6a1209e;
using Go = import "go.capnp";
$Go.package("proto");
$Go.import("github.com/scionproto/scion/go/proto");

struct RoutingPolicyExt{
    set @0 :Bool;   # Is the extension present? Every extension must include this field.
    polType @1 :UInt8;  # The policy type
    ifID @2 :UInt64;
    isdases @3 :List(UInt64);
}

struct ISDAnnouncementExt{
    set @0 :Bool;   # TODO(Sezer): Implement announcement extension
}

struct HiddenPathSegExtn{
    set @0 :Bool;
}

struct LinkmetricsExtn {
    struct Latencyinfo {
        latencynonpeeringclusters @0 :List(Lnpcluster);
        latencypeeringclusters @1 :List(Lpcluster);
        egresslatency @2 :UInt16;
        intooutlatency @3 :UInt16;

        struct Lnpcluster {
            clusterdelay @0 :UInt16;
            interfaces @1 :List(UInt16);
        }

        struct Lpcluster {
            clusterdelay @0 :UInt16;
            latencyinterfacepairs @1 :List(Lppair);

            struct Lppair {
                interface @0 :UInt16;
                interdelay @1 :UInt16;
            }
        }
    }
}
