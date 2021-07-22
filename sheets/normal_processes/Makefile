# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make normal_processes.pdf

normal_processes.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > normal_processes.tex

normal_processes.pdf: ../../*.tex ../../trademark.pdf *.tex normal_processes.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode normal_processes.tex
	../../latexrun normal_processes.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: normal_processes.pdf
	open normal_processes.pdf

o:
	make open

terms: normal_processes.tex
	bbk_sheet -mode ts normal_processes.tex

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
	rm -f normal_processes.*
	make reset
	bbk_sheet -mode mk > Makefile