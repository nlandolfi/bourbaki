# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make outcome_variable_events.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

outcome_variable_events.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > outcome_variable_events.tex

outcome_variable_events.pdf: ../../*.tex ../../trademark.pdf *.tex outcome_variable_events.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode outcome_variable_events.tex
	../../latexrun outcome_variable_events.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: outcome_variable_events.pdf
	open outcome_variable_events.pdf

o:
	make open

terms: outcome_variable_events.tex
	bbk_sheet -mode ts outcome_variable_events.tex

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
	rm -f outcome_variable_events.*
	make reset
	bbk_sheet -mode mk > Makefile