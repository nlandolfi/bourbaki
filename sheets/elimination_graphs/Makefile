# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make elimination_graphs.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

elimination_graphs.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > elimination_graphs.tex

elimination_graphs.pdf: ../../*.tex ../../trademark.pdf *.tex elimination_graphs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode elimination_graphs.tex
	../../latexrun elimination_graphs.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: elimination_graphs.pdf
	open elimination_graphs.pdf

o:
	make open

terms: elimination_graphs.tex
	bbk sheet -mode ts elimination_graphs.tex

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
	rm -f elimination_graphs.*
	make reset
	bbk sheet -mode mk > Makefile