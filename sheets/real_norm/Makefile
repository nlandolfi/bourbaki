# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_norm.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

real_norm.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_norm.tex

real_norm.pdf: ../../*.tex ../../trademark.pdf *.tex real_norm.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_norm.tex
	../../latexrun real_norm.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_norm.pdf
	open real_norm.pdf

o:
	make open

terms: real_norm.tex
	bbk_sheet -mode ts real_norm.tex

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
	rm -f real_norm.*
	make reset
	bbk_sheet -mode mk > Makefile