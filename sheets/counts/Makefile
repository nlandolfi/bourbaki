# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make counts.pdf

counts.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > counts.tex

counts.pdf: ../../*.tex ../../trademark.pdf *.tex counts.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode counts.tex
	../../latexrun counts.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: counts.pdf
	open counts.pdf

o:
	make open

terms: counts.tex
	bbk_sheet -mode ts counts.tex

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
	rm -f counts.*
	make reset
	bbk_sheet -mode mk > Makefile