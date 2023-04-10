# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make rational_order.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

rational_order.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > rational_order.tex

rational_order.pdf: ../../*.tex ../../trademark.pdf *.tex rational_order.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode rational_order.tex
	../../latexrun rational_order.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: rational_order.pdf
	open rational_order.pdf

o:
	make open

terms: rational_order.tex
	bbk sheet -mode ts rational_order.tex

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
	rm -f rational_order.*
	make reset
	bbk sheet -mode mk > Makefile