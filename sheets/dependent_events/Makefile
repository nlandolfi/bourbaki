# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make dependent_events.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

dependent_events.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > dependent_events.tex

dependent_events.pdf: ../../*.tex ../../trademark.pdf *.tex dependent_events.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode dependent_events.tex
	../../latexrun dependent_events.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: dependent_events.pdf
	open dependent_events.pdf

o:
	make open

terms: dependent_events.tex
	bbk_sheet -mode ts dependent_events.tex

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
	rm -f dependent_events.*
	make reset
	bbk_sheet -mode mk > Makefile