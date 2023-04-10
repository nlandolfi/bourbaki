# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make directed_acyclic_graphs.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

directed_acyclic_graphs.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > directed_acyclic_graphs.tex

directed_acyclic_graphs.pdf: ../../*.tex ../../trademark.pdf *.tex directed_acyclic_graphs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode directed_acyclic_graphs.tex
	../../latexrun directed_acyclic_graphs.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: directed_acyclic_graphs.pdf
	open directed_acyclic_graphs.pdf

o:
	make open

terms: directed_acyclic_graphs.tex
	bbk sheet -mode ts directed_acyclic_graphs.tex

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
	rm -f directed_acyclic_graphs.*
	make reset
	bbk sheet -mode mk > Makefile