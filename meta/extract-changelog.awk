#!/usr/bin/awk -f

BEGIN {
    release = target_release;
    in_section = 0;
}

/^## \[/ {
    match($0, /\[([^\]]+)\]/, ver);
    if (ver[1] == release) {
        in_section = 1;
    } else {
        in_section = 0;
    }
}

in_section && !/^## \[/ {
    if ($0 !~ /^[[:space:]]*$/) print $0;
}

