# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make product_manifolds.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

product_manifolds.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > product_manifolds.tex

product_manifolds.pdf: ../../*.tex ../../trademark.pdf *.tex product_manifolds.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode product_manifolds.tex
	../../latexrun product_manifolds.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: product_manifolds.pdf
	open product_manifolds.pdf

o:
	make open

terms: product_manifolds.tex
	bbk sheet -mode ts product_manifolds.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f product_manifolds.*
	make reset
	bbk sheet -mode mk > Makefile