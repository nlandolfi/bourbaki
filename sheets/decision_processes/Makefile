# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make decision_processes.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

decision_processes.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > decision_processes.tex

decision_processes.pdf: ../../*.tex ../../trademark.pdf *.tex decision_processes.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode decision_processes.tex
	../../latexrun decision_processes.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: decision_processes.pdf
	open decision_processes.pdf

o:
	make open

terms: decision_processes.tex
	bbk sheet -mode ts decision_processes.tex

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
	rm -f decision_processes.*
	make reset
	bbk sheet -mode mk > Makefile