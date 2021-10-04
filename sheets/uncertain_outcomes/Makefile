# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make uncertain_outcomes.pdf

uncertain_outcomes.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > uncertain_outcomes.tex

uncertain_outcomes.pdf: ../../*.tex ../../trademark.pdf *.tex uncertain_outcomes.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode uncertain_outcomes.tex
	../../latexrun uncertain_outcomes.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: uncertain_outcomes.pdf
	open uncertain_outcomes.pdf

o:
	make open

terms: uncertain_outcomes.tex
	bbk_sheet -mode ts uncertain_outcomes.tex

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
	rm -f uncertain_outcomes.*
	make reset
	bbk_sheet -mode mk > Makefile