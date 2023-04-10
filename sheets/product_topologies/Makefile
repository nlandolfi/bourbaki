# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make product_topologies.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

product_topologies.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > product_topologies.tex

product_topologies.pdf: ../../*.tex ../../trademark.pdf *.tex product_topologies.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode product_topologies.tex
	../../latexrun product_topologies.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: product_topologies.pdf
	open product_topologies.pdf

o:
	make open

terms: product_topologies.tex
	bbk sheet -mode ts product_topologies.tex

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
	rm -f product_topologies.*
	make reset
	bbk sheet -mode mk > Makefile