# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_set_closures.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

real_set_closures.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_set_closures.tex

real_set_closures.pdf: ../../*.tex ../../trademark.pdf *.tex real_set_closures.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_set_closures.tex
	../../latexrun real_set_closures.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_set_closures.pdf
	open real_set_closures.pdf

o:
	make open

terms: real_set_closures.tex
	bbk_sheet -mode ts real_set_closures.tex

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
	rm -f real_set_closures.*
	make reset
	bbk_sheet -mode mk > Makefile