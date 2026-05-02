# Maintainer: voidVoy7 <voidvoy7@proton.me>
pkgname=gtkcalc
pkgver=1.0
pkgrel=0
pkgdesc="A simple calculator made in Go with GTK"
arch=("x86_64")
url="https://github.com/voidVoy7/GTKCalc"
license=('GPL-3.0')
depends=('gtk4')
makedepends=('go')
source=("https://github.com/voidVoy7/GTKCalc/archive/refs/tags/${pkgver}.tar.gz")
sha256sums=('86f8cdd3a6ca67077ba8d3def1bc87118927437108abe75aab85dad9d5a11c7d')

build() {
    cd "GTKCalc-$pkgver"
    export CGO_CFLAGS="-Wno-builtin-declaration-mismatch"
    go build -buildmode=pie -trimpath -o gtkcalc
}

package() {
    cd "GTKCalc-$pkgver"
    install -Dm0755 gtkcalc "$pkgdir/usr/bin/gtkcalc"
    install -Dm0755 gtkcalc.desktop "$pkgdir/usr/share/applications/gtkcalc.desktop"
}
